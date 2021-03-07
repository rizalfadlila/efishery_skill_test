fs = require("fs");
crypto = require("crypto");

module.exports = {
  signup(req, res) {
    let password = crypto.randomBytes(2).toString("hex");
    let content = `${req.body.name},${req.body.phone},${req.body.role},${password}\n`;

    fs.readFile("./register.txt", function (err, data) {
      if (err) {
        return res.status(200).send({
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
              return res.status(200).send({
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
    return res.status(200).send({
      status: "signin",
    });
  },
};
