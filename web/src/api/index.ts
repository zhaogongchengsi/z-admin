import { Get } from "../utils/service";


export const verify = () => {
    return Get(`other/verify`);
}