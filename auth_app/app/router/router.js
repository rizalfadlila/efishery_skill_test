const authService = require("../services").authService;
const userService = require("../services").userService;
const mwAuth = require("../middlewares").auth;

module.exports = function (app) {
  app.post("/signup", (req, res) => {
    /*  #swagger.parameters['object'] = {
                in: 'body',
                required: true,
                schema: {
                  "name": "admin",
                  "phone": "12345",
                  "role": "admin"
                }
        },
        #swagger.responses[401] = { description: 'Unauthorized' },
        #swagger.responses[200] = { description: 'Ok' },
        #swagger.responses[500] = { description: 'Internal Server Error' },
  */

    authService.signup(req, res);
  });
  app.post("/signin", (req, res) => {
    /*  #swagger.parameters['object'] = {
                in: 'body',
                required: true,
                schema: {
                  "username": "admin",
                  "password": "12345"
                }
        },
        #swagger.responses[401] = { description: 'Unauthorized' },
        #swagger.responses[200] = { description: 'Ok' },
        #swagger.responses[500] = { description: 'Internal Server Error' },
  */
    authService.signin(req, res);
  });

  app.get("/claims-jwt", [mwAuth.verifyToken], (req, res) => {
    /*  #swagger.parameters['Authorization'] = {
                in: 'header',
                required: true,
                description: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJwaG9uZSI6IjEyMzQ1Iiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMS0wMy0wOFQwOTozMDoyOS41OTBaIiwiaWF0IjoxNjE1MTk2MjM5LCJleHAiOjE2MTUyODI2Mzl9.mEIZHstVyG4VBga3x3ONzks0B8Pv7JZ5QWFuC44Jsxo'
        },
        #swagger.responses[401] = { description: 'Unauthorized' },
        #swagger.responses[200] = { description: 'Ok' },
        #swagger.responses[500] = { description: 'Internal Server Error' },
  */
    userService.claimsJwt(req, res);
  });
};
