import { createServer } from "miragejs";



export function mirage(){
    createServer({
        routes() {
          this.get("/api/groups/:departamentId", (_, request) => {
            const departamentId = +request.params.departamentId
            return { groups: [
              { id: 1, title: "Учителя", message: "T.", departamentId: 1},
              { id: 2, title: "267", message: "B.", departamentId: 2},
              { id: 3, title: "265", message: "Crain",  departamentId: 3},
              { id: 4, title: "129", message: "T.", departamentId: 1},
              { id: 5, title: "340", message: "B.", departamentId: 2},
              { id: 6, title: "450", message: "Crain",  departamentId: 3}
          ].filter(group => group.departamentId === departamentId)}
          })

          this.get("/api/depataments/", () => ({
            groups: [
                { id: 1, title: "АИВТ", number: 12 },
                { id: 2, title: "Автосервис", number: 12 },
                { id: 3, title: "Технология и дизайн", number: 12 }
            ],
          }))
        },
      })
 }

