<template>
  <div class="login-container">
    <div class="login-center">
      <div class="login-x login-img">
        <img src="/z-admin.png" alt="z-admin" />
      </div>
      <div class="login-x login-from">
        <div class="from">
          <el-form label-position="top" :model="LoginForm" :rules="rules" ref="fromRef" size="large" >
            <el-form-item label="账号" :rules="rules.username" prop="username">
              <el-input v-model="LoginForm.username" placeholder="请输入账号"></el-input>
            </el-form-item>
            <el-form-item label="密码" :rules="rules.password" prop="password">
              <el-input v-model="LoginForm.password" placeholder="请输入密码"></el-input>
            </el-form-item>
            <el-form-item label="验证码" :rules="rules.verify" prop="verify">
              <div class="verifi-input" >
                <el-input v-model="LoginForm.verify" placeholder="请输入验证码"></el-input>
                <div class="verifi" @click="getVerify">
                  <img :src="codeImg.code" alt="">
                </div>
              </div>
            </el-form-item>
          </el-form>
          <el-button type="primary" size="large" @click="submitForm(fromRef)">登录</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import "./login.scss";
import { reactive, ref, onMounted, } from "vue";
import { Login, rules } from "./login";
import { FormInstance } from "element-plus";
import { verify, login } from "@/api/login";
import { userStore } from "@/pinia/index";
import { useRouter } from "vue-router";

const user = userStore()
const router = useRouter()


const LoginForm = reactive<Login>({
  username: "admin",
  password: "abc123",
  verify: "",
});

const codeImg = reactive({
  id: "",
  code: "",
})

const fromRef = ref<FormInstance>();

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      const data:Login  = {
        verify_Id: codeImg.id,
        ...LoginForm,
      }
      const isLogin = await user.getUser(data)
      if (isLogin) {
        router.push("/")
      } else {
        getVerify()
      }
    } else {
      getVerify()
    }
  })
};


const getVerify = async () => {
  try {
      const codeData = await verify()
      const { id, code } = codeData.data
      codeImg.id = id
      codeImg.code = code
  } catch (e) {
    console.log(e)
  }

}
onMounted(getVerify);

</script>
