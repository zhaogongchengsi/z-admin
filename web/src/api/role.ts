import { Get } from "../utils/service";



export function role () {
    return Get("/role/index")
}