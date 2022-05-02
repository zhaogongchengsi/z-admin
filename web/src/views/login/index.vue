<template>
  <div class="login-container">
    <div class="login-center">
      <div class="login-x login-img">这里放图片</div>
      <div class="login-x login-from">
        <div class="from">
          <el-form label-position="top">
            <el-form-item label="账号" required>
              <el-input placeholder="请输入账号"></el-input>
            </el-form-item>
            <el-form-item label="密码" required>
              <el-input placeholder="请输入密码"></el-input>
            </el-form-item>
            <el-form-item label="验证码" required>
              <div class="verifi-input">
                <el-input placeholder="请输入验证码"></el-input>
                <div class="verifi">验证码</div>
              </div>
            </el-form-item>
          </el-form>
          <el-button type="primary" size="large" @click="handleLogin">登录</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import "./login.scss";
import { ref } from "vue";

const handleLogin = () => {};
</script>
