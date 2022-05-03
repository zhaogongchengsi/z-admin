import { ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { login } from '../api/login'
import { Login } from '@/views/login/login'

export const userStore = defineStore("user", {
    state: () => {
        return {
            user: {},
            token: null,
        }
    },

    getters: {

    },

    actions: {
       async getUser (loginData:Login) {
        try {
            const { code, data, msg }  = await login<Login>(loginData)

            if (code === 200) {
                // @ts-ignore
                this.user = data.user
                // @ts-ignore
                this.token = data.z_token
                localStorage.setItem('z_token', data.z_token)
                ElMessage.success('登录成功')
                return true
            } else {
                ElMessage.error(msg)
                return false
            }
        } catch (err) {
            ElMessage.error("登录失败,请重试")
            console.error(err)
            return false
        }

       }
    },
})