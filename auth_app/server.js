const express = require("express");
const app = express();
const swaggerUi = require("swagger-ui-express");
const swaggerFile = require("./swagger_output.json");

app.use("/doc", swaggerUi.serve, swaggerUi.setup(swaggerFile));

require("./app/middlewares/middleware.js")(app);
require("./app/router/router.js")(app);

app.listen(process.env.PORT, () =>
  console.log(
    "Server is running!\nAPI documentation: http://localhost:3000/doc"
  )
);
