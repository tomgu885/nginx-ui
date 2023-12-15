<script setup lang="ts">
// import {}
import {computed, ref} from "vue";

const props = defineProps(['value', 'mask'])
const emit = defineEmits(['update:value'])


const options = computed(() => {
    const _options = []
    for (const [key, value] of Object.entries(props.mask)) {
        // let v = (typeof value == 'string') value : (value as ()=> string) value())
        let v
        if (typeof value == 'function') {
            v = value()
        } else {
            v = value
        }
        console.log('typeof key', typeof key, '| key:', key)
        _options.push({label: v, value: key})
    }

    return _options
})

const _value = computed({
    get() {
        let v
        if (typeof props.mask?.[props.value] == "function") {
            v = (props.mask?.[props.value] as () => string)()
        } else if (typeof props.mask?.[props.value] == 'string') {
            v = props.mask[props.value]
        } else {
            v = props.value
        }

        return v
    },
    set (v) {
        console.log('radio update:value', v);
        emit('update:value', v)
    }
})

</script>
<template>
    <a-radio-group v-model:value="_value" :options="options" button-style="solid"  />
</template>
