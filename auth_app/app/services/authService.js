const fs = require("fs");
const crypto = require("crypto");
const jwt = require("jsonwebtoken");
const filename = "./register.txt";

var self = (module.exports = {
  signup(req, res) {
    let password = crypto.randomBytes(2).toString("hex");
    const now = new Date().toISOString();
    let content = `${req.body.name},${password},${req.body.phone},${req.body.role},${now}\n`;

    self.writeFile(req.body.name, content, res);
  },

  signin(req, res) {
    fs.readFile(filename, "utf8", function (err, data) {
      if (err) {
        return res.status(401).send({
          message: "Invalid username or password",
        });
      }
      var found =
        data.indexOf(`${req.body.username},${req.body.password}`) >= 0;

      if (found) {
        return res.status(200).send({
          accessToken: self.generateToken(data),
        });
      } else {
        return res.status(401).send({
          message: "Invalid username or password",
        });
      }
    });
  },

  generateToken: function (data) {
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

    return token;
  },

  writeFile: function (query, content, res) {
    fs.readFile(filename, function (err, data) {
      if (err) {
        self.appendRecord(content, res);
      } else {
        if (data.indexOf(query) >= 0) {
          return res.status(409).send({
            message: "Username already exists",
          });
        } else {
          self.appendRecord(content, res);
        }
      }
    });
  },
  appendRecord: function (content, res) {
    fs.appendFile(filename, content, { encoding: "utf8" }, function (err) {
      if (err) {
        return res.status(500).send({
          error: err,
        });
      }
      return res.status(200).send({
        message: "Registered successfully",
      });
    });
  },
});
