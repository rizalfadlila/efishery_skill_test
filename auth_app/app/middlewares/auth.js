const jwt = require("jsonwebtoken");

module.exports = {
  verifyToken(req, res, next) {
    var tokenHeader = "";

    if (req.headers["authorization"] != "") {
      tokenHeader = req.headers["authorization"];
    }

    if (tokenHeader.split(" ")[0] !== "Bearer") {
      return res.status(401).send({
        message: "Unauthorized",
        errors: "Incorrect token format",
      });
    }

    let token = tokenHeader.split(" ")[1];

    if (!token) {
      return res.status(401).send({
        message: "Unauthorized",
        errors: "No token provided",
      });
    }

    jwt.verify(token, process.env.APP_KEY, (err, decoded) => {
      if (err) {
        return res.status(401).send({
          message: "Unauthorized",
          errors: err,
        });
      }

      req.name = decoded.name;
      req.phone = decoded.phone;
      req.role = decoded.role;
      req.timestamp = decoded.timestamp;

      next();
    });
  },
};
