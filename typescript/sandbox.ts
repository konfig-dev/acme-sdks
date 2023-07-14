import { Acme } from "./index";

const acme = new Acme({ apiKey: "MY_API_KEY" });
async function main() {
  const pet = await acme.pet.add({
    name: "Dog",
    photoUrls: ["myphoto.com"],
  });
  const fetched = await acme.pet.getById({ petId: 20 });
}
