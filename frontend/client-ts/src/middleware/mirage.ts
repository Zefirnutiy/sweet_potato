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

          this.get("/api/users/:groupId", (_, request) => {
            const groupId = +request.params.groupId
            return { users: [
              { id: 1, firstName: "Сергей Орлов", message: "T.", groupId: 1},
              { id: 1, firstName: "Сергей Орлов", message: "T.", groupId: 1},
              { id: 1, firstName: "Сергей Орлов", message: "T.", groupId: 1},
              { id: 2, firstName: "Василий Попов", message: "B.", groupId: 2},
              { id: 2, firstName: "Василий Попов", message: "B.", groupId: 2},
              { id: 2, firstName: "Василий Попов", message: "B.", groupId: 2},
              { id: 2, firstName: "Алексей Попов", message: "B.", groupId: 3},
              { id: 2, firstName: "Алексей Попов", message: "B.", groupId: 3},
              { id: 2, firstName: "Максим Попов", message: "B.", groupId: 4},
              { id: 2, firstName: "Илья Попов", message: "B.", groupId: 5},
              { id: 2, firstName: "Мария Попова", message: "B.", groupId: 6},
          ].filter(user => user.groupId === groupId)}
          })

          this.get("/api/depataments/", () => ({
            groups: [
                { id: 1, title: "АИВТ", message: "12 групп" },
                { id: 2, title: "Автосервис", message: "14 групп" },
                { id: 3, title: "Технология и дизайн", message: "15 групп" },
            ],
          }))
        },
      })
 }

