const authService = require("../services").authService;
const userService = require("../services").userService;

module.exports = function (app) {
  app.post("/signup", authService.signup);
  app.post("/signin", authService.signin);

  app.get("/claims-jwt", userService.claimsJwt);
};
