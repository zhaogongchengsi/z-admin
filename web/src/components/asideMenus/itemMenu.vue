<template>
  <el-menu-item :index="index">
    <el-icon>
      <component :is="menuInfo.icon" />
    </el-icon>
    <span>{{menuInfo.label}}</span>
  </el-menu-item>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "AsideItemMenu",
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
    index: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    console.log(props.index);
  },
});
</script>