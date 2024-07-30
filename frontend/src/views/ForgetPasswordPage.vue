<script setup>
import { nextTick, ref } from 'vue'

const step = ref(0)

const steps = [
  { label: 'Отправка письма' },
  { label: 'Получение кода', completed:false },
  { label: 'Изменение пароля' },
  { label: 'Готово!'},
]

let finished = () => {
    window.location.href = '/sign_in'
}
</script>

<template>
    <div class="relative main h-screen w-full">
        <div class="absolute shadow-xl bg-white p-5 w-1/2 left-1/4 top-1/3 rounded-xl mx-auto">
            <VaStepper v-model="step" @finish="finished()" :steps="steps">
                <template #step-content-0>
                    <form class="">
                        <h1 class="text-3xl">Восстановление пароля</h1>
                        <div class="mt-4">
                            <label class="block mb-2 text-xl text-gray-600 ">Почта</label>
                            <input type="email" placeholder="Введите почту" class="text-xl block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                        </div>
                    </form>
                </template>
                <template #step-content-1>
                    <form class="">
                        <h1 class="text-3xl">Введите код который пришёл на почту Adamanq@yandex.ru</h1>
                        <div class="mt-4">
                            <label class="block mb-2 text-xl text-gray-600 ">Код</label>
                            <input type="text" maxlength="5" placeholder="Введите почту" class="block text-xl w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                        </div>
                    </form>
                </template>
                <template #step-content-2>
                    <form class="">
                        <h1 class="text-3xl">Введите новый пароль</h1>
                        <div class="mt-4">
                            <label class="block mb-2 text-xl text-gray-600 ">Почта</label>
                            <input type="password" placeholder="Введите пароль" class="text-xl block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                        </div>
                    </form>
                </template>
                <template #step-content-3>
                    <h1>Пароль был успешно изменён!</h1>
                    <button>CLickl</button>
                </template>
            </VaStepper>
        </div>
    </div>
</template>

<style scoped>
    .main {
        background-image: url('https://sun9-80.userapi.com/impg/kPr3--WrF0ZlbTM6wyiolQRNzXKsjkVQolLzAw/ejfqKE1l7z0.jpg?size=2048x1366&quality=95&sign=388fba0f019f73ec2663375b05f5e475&type=album'); 
        background-size: 100% 100%; 
        background-repeat: no-repeat; 
        background-position: center;
    }
</style>