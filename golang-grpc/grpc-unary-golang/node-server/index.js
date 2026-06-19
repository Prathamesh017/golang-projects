const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const packageDefinition = protoLoader.loadSync(
  "../proto/user.proto"
);

const proto = grpc.loadPackageDefinition(
  packageDefinition
);

function main(){
const client = new proto.user.UserService(
  "localhost:8081",
  grpc.credentials.createInsecure()
);
client.GetUser(
  { id: 1 },
  (err, response) => {

     console.log(response)

  }
)
}

main();