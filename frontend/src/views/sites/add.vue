<script setup lang="ts">
import gettext from '@/gettext'
const {$gettext} = gettext
function add() {

}

</script>

<template>
    <a-card title="添加">
        <a @click="add"></a>
        <div class="domain-add-container">
            <a-step></a-step>
        </div>
    </a-card>
</template>

<style scoped lang="less">

</style>
