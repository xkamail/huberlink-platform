import axios, {AxiosResponse} from "axios"

const apiURL = process.env.NEXT_PUBLIC_API_URL

const instance = axios.create({
    baseURL: apiURL,
    timeout: 10 * 1000,
})

instance.interceptors.request.use(config => {
    return config
})

instance.interceptors.response.use((f: AxiosResponse) => {
    return f
}, (error) => {
    return Promise.reject(error)
})