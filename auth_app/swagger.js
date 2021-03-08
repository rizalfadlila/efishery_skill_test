const swaggerAutogen = require("swagger-autogen")();

const doc = {
  info: {
    version: "1.0.0",
    title: "Auth App",
    description: "Documentation Auth App.",
  },
  host: "localhost:3000",
  basePath: "/",
  schemes: ["http"],
  consumes: ["application/json"],
  produces: ["application/json"],
};

const outputFile = "./swagger_output.json";
const endpointsFiles = ["./app/router/router.js"];

swaggerAutogen(outputFile, endpointsFiles, doc).then(() => {
  require("./server"); // Your project's root file
});
