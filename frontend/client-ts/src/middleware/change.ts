export function changeSchema(){
    document.addEventListener('keydown', function(event) {
      if (event.code === 'KeyZ') {
        document.documentElement.style.setProperty("--main-hue", "250" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      } 
  
      if (event.code === 'KeyX') {
        document.documentElement.style.setProperty("--main-hue", "280" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      } 
  
      if (event.code === 'KeyC') {
        document.documentElement.style.setProperty("--main-hue", "310" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      } 
  
      if (event.code === 'KeyV') {
        document.documentElement.style.setProperty("--main-hue", "340" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      }

      if (event.code === 'KeyB') {
        document.documentElement.style.setProperty("--main-hue", "370" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      }

      if (event.code === 'KeyN') {
        document.documentElement.style.setProperty("--main-hue", "0" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      }
    });
  }

export function changePages(setPage: any){
   document.addEventListener('keydown', function(event) {
    if(event.code === 'Digit1'){
      setPage("admin")
      return
    }

    if(event.code === 'Digit2'){
      setPage("user")
      return 
    }
    
    if(event.code === 'Digit3')
      setPage("register") 
  })
}