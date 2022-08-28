export function changeSchema(){
    document.addEventListener('keydown', function(event) {
      if (event.code === 'KeyZ') {
        document.documentElement.style.setProperty("--main-hue", "250" )
        document.documentElement.style.setProperty("--accent-hue", "40" )
        return
      } 
  
      if (event.code === 'KeyX') {
        document.documentElement.style.setProperty("--main-hue", "50" )
        document.documentElement.style.setProperty("--accent-hue", "160" )
        return
      } 
  
      if (event.code === 'KeyC') {
        document.documentElement.style.setProperty("--main-hue", "150" )
        document.documentElement.style.setProperty("--accent-hue", "240" )
        return
      } 
  
      if (event.code === 'KeyV') {
        document.documentElement.style.setProperty("--main-hue", "200" )
        document.documentElement.style.setProperty("--accent-hue", "300" )
        return
      }
    });
  }