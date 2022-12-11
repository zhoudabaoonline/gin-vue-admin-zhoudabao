
<template>
    <div>
        <el-upload v-model:file-list="imageinit" :action="`${path}/fileUploadAndDownload/upload`"
            :headers="{ 'x-token': userStore.token }" list-type="picture-card" :on-success="handleImageSuccess"
            :before-upload="beforeImageUpload" :multiple="multiple" :auto-upload="true" :on-remove="handleRemove"
            :on-preview="handlePictureCardPreview">
            <el-icon>
                <Plus />
            </el-icon>
        </el-upload>
        <el-dialog v-model="dialogVisible">
            <img w-full :src="dialogImageUrl" alt="Preview Image" />
        </el-dialog>
    </div>
</template>
  
<script setup>
import ImageCompress from '@/utils/image'
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { Plus } from '@element-plus/icons-vue'

const emit = defineEmits(['on-success'])
const props = defineProps({
    imageUrl: {
        type: String,
        default: ''
    },
    fileSize: {
        type: Number,
        default: 2048 // 2M 超出后执行压缩
    },
    maxWH: {
        type: Number,
        default: 1920 // 图片长宽上限
    },
    multiple: {
        type: Boolean,
        default: false // 是否多图上传
    },
    limit: {
        type: Number,
        default: 1920 // 图片长宽上限
    }
})

const path = ref(import.meta.env.VITE_BASE_API)
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const fileList = ref([])
const imageinit = ref([])
const userStore = useUserStore()

onMounted(() => {
    if (props.imageUrl != '') {
        let t = JSON.parse(props.imageUrl)
        for (let v in t) {
            t[v].url = path + "/" + t[v].url
        }
        imageinit.value = t
    } else {
        imageinit.value = []
    }
})

const beforeImageUpload = (file) => {
    const isJPG = file.type === 'image/jpeg'
    const isPng = file.type === 'image/png'
    if (!isJPG && !isPng) {
        ElMessage.error('上传头像图片只能是 jpg或png 格式!')
        return false
    }

    const isRightSize = file.size / 1024 < props.fileSize
    if (!isRightSize) {
        // 压缩
        const compress = new ImageCompress(file, props.fileSize, props.maxWH)
        return compress.compress()
    }
    return isRightSize
}

const handleImageSuccess = (res, uploadFile, uploadFiles) => {
    let t = []
    for (const o of uploadFiles) {
        if (o.response != undefined) {
            t.push({ name: o.response.data.file.name, key: o.response.data.file.key, url: o.response.data.file.url })
        } else {
            t.push(o)
        }
    }
    emit('on-success', JSON.stringify(t))
}



const handleRemove = (uploadFile, uploadFiles) => {
    let t = []
    for (const o of uploadFiles) {
        if (o.response != undefined) {
            t.push({ name: o.response.data.file.name, key: o.response.data.file.key, url: o.response.data.file.url })
        } else {
            t.push(o)
        }
    }
    emit('on-success', JSON.stringify(t))

}

const handlePictureCardPreview = (uploadFile) => {
    dialogImageUrl.value = uploadFile.url
    dialogVisible.value = true
}




</script>
  
<script>

export default {
    name: 'UploadImageSelf',
    methods: {

    }
}
</script>
  
<style lang="scss" scoped>

</style>
  