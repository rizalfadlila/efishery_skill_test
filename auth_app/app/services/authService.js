const fs = require("fs");
const crypto = require("crypto");
const jwt = require("jsonwebtoken");

module.exports = {
  signup(req, res) {
    let password = crypto.randomBytes(2).toString("hex");
    const now = new Date().toISOString();
    let content = `${req.body.name},${password},${req.body.phone},${req.body.role},${now}\n`;

    fs.readFile("./register.txt", function (err, data) {
      if (err) {
        return res.status(500).send({
          error: err,
        });
      }
      if (data.indexOf(req.body.name) >= 0) {
        return res.status(200).send({
          error: "Username already exists",
        });
      } else {
        fs.appendFile(
          "register.txt",
          content,
          { encoding: "utf8" },
          function (err) {
            if (err)
              return res.status(500).send({
                error: err,
              });
          }
        );
        return res.status(200).send({
          message: "Registered Successfully",
        });
      }
    });
  },

  signin(req, res) {
    fs.readFile("./register.txt", "utf8", function (err, data) {
      if (err) {
        return res.status(200).send({
          error: err,
        });
      }
      if (data.indexOf(`${req.body.username},${req.body.password}`) >= 0) {
        const records = data.replace("\n", "").split(",");
        var payload = {
          name: records[0],
          phone: records[2],
          role: records[3],
          timestamp: records[4],
        };
        var token = jwt.sign(payload, process.env.APP_KEY, {
          expiresIn: 86400,
          algorithm: "HS256",
        });
        return res.status(200).send({
          accessToken: token,
        });
      } else {
        return res.status(401).send({
          message: "Invalid username or password",
        });
      }
    });
  },
};
