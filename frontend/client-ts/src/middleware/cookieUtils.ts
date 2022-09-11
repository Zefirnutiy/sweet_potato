export function setCookie ( name: string, value: string ){
  var cookie_string = name + "=" + value;
  document.cookie = cookie_string;
}

export function deleteCookie (name: string)
{
  var cookie_date = new Date ( );  
  cookie_date.setTime ( cookie_date.getTime() - 1 );
  document.cookie = name += "=; expires=" + cookie_date.toUTCString();
}

export function getCookie (name: string)
{
  var results = document.cookie.match ( '(^|;) ?' + name + '=([^;]*)(;|$)' );
 
  if ( results )
    return (  decodeURIComponent( results[2] ) );
  else
    return "";
}