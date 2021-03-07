const authService = require("../services").authService;
const userService = require("../services").userService;
const mwAuth = require("../middlewares").auth;

module.exports = function (app) {
  app.post("/signup", authService.signup);
  app.post("/signin", authService.signin);

  app.get("/claims-jwt", [mwAuth.verifyToken], userService.claimsJwt);
};
