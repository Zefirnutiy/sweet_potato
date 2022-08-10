import {useState, useCallback, useEffect} from "react"
import { setCookie, getCookie } from "../middleware/cookieUtils"


export const useAuth = () => {

	
	const [token, setToken] = useState<string | null>(null)
	const [userId, setUserId] = useState<number | null>(null)
	const [ready, setReady] = useState<boolean>(false)


	const login = useCallback((jwtToken: string, id: number) => {
		setToken(jwtToken)
		setUserId(id)

		setCookie('userData', JSON.stringify({
			userId: id, token: jwtToken
		}))

	}, [])


	const logout = useCallback(() => {
		setToken(null)
		setUserId(null)
		localStorage.removeItem('userData')

	}, [])

	useEffect(() => {
        const userData = getCookie('userData')

		const data = JSON.parse(userData || "")

		if(data && data.token){
			login(data.token, data.userId)
		}
		setReady(true)

	}, [login])


	return {logout, login, token, userId, ready}
}