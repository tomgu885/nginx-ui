<script setup lang="ts">
import gettext from '@/gettext'
const {$gettext} = gettext
function add() {

}

</script>

<template>
    <a-card title="添加">
        <a @click="add">1</a>
        <div class="domain-add-container">
            <a-step>1</a-step>
            <a-step>2</a-step>
        </div>
    </a-card>
</template>

<style scoped lang="less">

</style>
