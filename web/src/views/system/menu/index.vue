
<template>
  <div class="menu-container">
    <div class="menu-header">
      <el-button type="primary" @click="openDialog">新增菜单</el-button>
    </div>

    <el-table :data="menuData" style="width: 100%" stripe>
      <el-table-column prop="name" label="菜单名称" width="180" />
      <el-table-column prop="path" label="路由路径" width="180" />
      <el-table-column prop="icon" label="图标" />
      <el-table-column prop="components" label="组件名称" />
      <el-table-column prop="parent_id" label="父菜单id" />
      <el-table-column prop="redirect" label="重定向路径" />
      <el-table-column prop="hidden" label="是否展示" />
    </el-table>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :before-close="handleClose">
      <el-form label-position="top" label-width="100px" :model="menuFrom">
        <el-form-item label="父菜单id">
          <el-input-number v-model="menuFrom.parent_id" :min="0" />
        </el-form-item>
        <el-form-item label="组件名称">
          <el-input v-model="menuFrom.name" clearable />
        </el-form-item>
        <el-form-item label="路由路径">
          <el-input v-model="menuFrom.path" clearable />
        </el-form-item>
        <el-form-item label="组件路径">
          <el-input v-model="menuFrom.components" clearable />
        </el-form-item>
        <el-form-item label="组件图标">
          <el-input v-model="menuFrom.icon" placeholder="请选择图标..." @focus="opendrawer">
            <template #prefix>
              <el-icon class="el-input__icon">
                <component :is="menuFrom.icon" />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="重定向路径">
          <el-input v-model="menuFrom.redirect" clearable />
        </el-form-item>
        <el-form-item label="是否出现">
          <el-switch v-model="menuFrom.hidden" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="createMenu">创建</el-button>
        </span>
      </template>
    </el-dialog>
    <el-drawer v-model="drawer" title="选择图标" :with-header="false" size="70%">
      <el-scrollbar>
        <div class="icon-list">
          <div
            v-for="(name, index) in IconNames"
            :key="index"
            class="icon-item"
            @click="chooseIcon(name)"
          >
            <el-tooltip effect="light" :content="name">
              <el-icon :size="30">
                <component :is="name" />
              </el-icon>
            </el-tooltip>
          </div>
        </div>
      </el-scrollbar>
    </el-drawer>
  </div>
</template>


<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { IconNames } from "@/plugin/icon";
import { Menu } from "./index";
import { addMenu, getMenuList } from "@/api/menu";

const menuFrom = reactive<Menu>({
  name: "",
  path: "",
  icon: "Apple",
  components: "",
  parent_id: 0,
  redirect: "",
  hidden: true,
});

const menuData = ref<Menu[]>([]);

onMounted(async () => {
  const res = await getMenuList();

  if (res.code === 200) {
    menuData.value = res.data;
  }
});

const dialogVisible = ref(false);
const dialogTitle = ref("新增菜单");

const handleClose = () => {
  dialogVisible.value = false;
};

const openDialog = () => {
  dialogVisible.value = true;
};

const drawer = ref(false);

const opendrawer = () => {
  drawer.value = true;
};

const chooseIcon = (name: string) => {
  menuFrom.icon = name;
  drawer.value = false;
};

const createMenu = async () => {
  const res = await addMenu(menuFrom);
  if (res.code === 200) {
    ElMessage.success("添加成功");
    dialogVisible.value = false;
  } else {
    ElMessage.error(res.msg);
  }
};
</script>


<style scoped lang="scss">
.menu-header {
  padding-bottom: 10px;
}

.menu-container {
  width: 100%;
}

.icon-list {
  display: grid;
  grid-template-columns: repeat(15, 1fr);
  grid-gap: 10px;
  grid-row-gap: 20px;
  .icon-item {
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    width: 35px;
    height: 35px;
    border-radius: 3px;
    &:hover {
      cursor: pointer;
      border: 1px solid var(--el-border-color);
    }
  }
}
</style>