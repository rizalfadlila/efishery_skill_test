module.exports = {
  signup(req, res) {
    return res.status(200).send({
      status: "signup",
    });
  },
  signin(req, res) {
    return res.status(200).send({
      status: "signin",
    });
  },
};
