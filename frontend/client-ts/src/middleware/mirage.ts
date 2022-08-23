import { createServer } from "miragejs";



export function mirage(){
    createServer({
        routes() {
          this.get("/api/groups/:DepartamentId", (_, request) => {
            const DepartamentId = +request.params.DepartamentId
            return { groups: [
              { Id: 1, Title: "Учителя", Message: "T.", DepartamentId: 1},
              { Id: 2, Title: "267", Message: "B.", DepartamentId: 2},
              { Id: 3, Title: "265", Message: "Crain",  DepartamentId: 3},
              { Id: 4, Title: "129", Message: "T.", DepartamentId: 1},
              { Id: 5, Title: "340", Message: "B.", DepartamentId: 2},
              { Id: 6, Title: "450", Message: "Crain",  DepartamentId: 3}
          ].filter(group => group.DepartamentId === DepartamentId)}
          })

          this.get("/api/users/:GroupId", (_, request) => {
            const GroupId = +request.params.GroupId
            return { users: [
              { Id: 1, FirstName: "Сергей Орлов", Message: "T.", GroupId: 1},
              { Id: 1, FirstName: "Сергей Орлов", Message: "T.", GroupId: 1},
              { Id: 1, FirstName: "Сергей Орлов", Message: "T.", GroupId: 1},
              { Id: 2, FirstName: "Василий Попов", Message: "B.", GroupId: 2},
              { Id: 2, FirstName: "Василий Попов", Message: "B.", GroupId: 2},
              { Id: 2, FirstName: "Василий Попов", Message: "B.", GroupId: 2},
              { Id: 2, FirstName: "Алексей Попов", Message: "B.", GroupId: 3},
              { Id: 2, FirstName: "Алексей Попов", Message: "B.", GroupId: 3},
              { Id: 2, FirstName: "Максим Попов", Message: "B.", GroupId: 4},
              { Id: 2, FirstName: "Илья Попов", Message: "B.", GroupId: 5},
              { Id: 2, FirstName: "Мария Попова", Message: "B.", GroupId: 6},
          ].filter(user => user.GroupId === GroupId)}
          })

          this.get("/api/depataments/", () => ({
            groups: [
                { Id: 1, Title: "АИВТ", Message: "12 групп" },
                { Id: 2, Title: "Автосервис", Message: "14 групп" },
                { Id: 3, Title: "Технология и дизайн", Message: "15 групп" },
            ],
          }))
        },
      })
 }

