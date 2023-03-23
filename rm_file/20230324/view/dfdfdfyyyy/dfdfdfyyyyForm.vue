<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所发生的:" prop="顶顶顶顶">
          <el-input v-model="formData.顶顶顶顶" :clearable="true" placeholder="请输入" />

  <!--自定义-->

<!--自定义结束-->
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dfdfdfyyyy'
}
</script>

<script setup>
import {
  createDfdfdfyyyy,
  updateDfdfdfyyyy,
  findDfdfdfyyyy
} from '@/api/dfdfdfyyyy'

import UploadImageSelf from '@/components/upload/imageself.vue'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { Delete, Download, Plus, ZoomIn } from '@element-plus/icons-vue'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            顶顶顶顶: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDfdfdfyyyy({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redfdfdfyyy
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDfdfdfyyyy(formData.value)
               break
             case 'update':
               res = await updateDfdfdfyyyy(formData.value)
               break
             default:
               res = await createDfdfdfyyyy(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
