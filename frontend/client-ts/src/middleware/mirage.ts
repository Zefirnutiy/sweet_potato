import { createServer } from "miragejs";



export function mirage(){
    createServer({
        routes() {
          this.get("/api/groups/", () => ({
            groups: [
                { id: 1, title: "Учителя", message: "T." },
                { id: 2, title: "267", message: "B." },
                { id: 3, title: "265", message: "Crain" }
            ],
          }))
        },
      })
 }

