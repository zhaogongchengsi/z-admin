<template>
  <div class="login-container">
    <div class="login-center">
      <div class="login-x login-img">这里放图片</div>
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
              <div class="verifi-input">
                <el-input v-model="LoginForm.verify" placeholder="请输入验证码"></el-input>
                <div class="verifi">
                  <img :src="code.code" alt="">
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
import { verify } from "@/api/index";

const LoginForm = reactive<Login>({
  username: "",
  password: "",
  verify: "",
});

const code = ref({
  id: "",
  code: "",
})

const fromRef = ref<FormInstance>();

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
    } else {
      console.log('error submit!', fields)
    }
  })
};


onMounted(async () => {
  const codeData = await verify()
  code.value = codeData.data
});

</script>
