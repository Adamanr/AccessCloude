<script setup>
import axios from 'axios';


let User = {
    email: "",
    password: "",
    remember: false,
}

async function sign_in() {
    try {
        axios.post("http://localhost:8080/users/sign_in", User)
        .then(response=> alert(response.statusText))
        .catch(error => alert(error))
    }catch(error){
        alert(error)
    }
   
}

// let sign_in = async () => {
    
//     // axios.post("http://localhost:8080/users/sign_in", User)
//     //     .then(response => alert(response))
//     //     .catch(error => alert(error))
//     axios.get("https://api.npms.io/v2/invalid-url")
//     .then(response => this.totalVuePackages = response.data.total)
//     .catch(error => {
//       this.errorMessage = error.message;
//       console.error("There was an error!", error);
//     });
// }
</script>


<template>
 <section class="bg-white dark:bg-[#f2f2f2]">
    <div class="flex h-full bg-[#f2f2f2] justify-center min-h-screen">

        <div class="flex items-center p-8 bg-[#f2f2f2] mx-auto lg:px-12 ">
            <div class="w-full">
                <h1 class="text-4xl font-semibold tracking-wider text-gray-800 capitalize ">
                    Вход в аккаунт
                </h1>

                <p class="mt-4 text-2xl text-gray-500 ">
                    Благодаря RedNews вы сможете получать самые актуальные новости на данный момент.<br>Если же у вас ещё нет аккаунта - <a href="/sign_up" class="text-blue-500">Регистрация</a> 
                </p>

                <form class="grid grid-cols-1 gap-6 mt-8 md:grid-cols-2" action="" method="get" target="_self">
            
                    <div class="mt-4">
                        <label class="block mb-2 text-xl text-gray-600 ">Почта</label>
                        <input v-model="User.email"  type="email" placeholder="Введите почту" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                    </div>

                    <div class="mt-4">
                        <label class="block mb-2 text-xl text-gray-600 ">Пароль</label>
                        <input type="password" placeholder="Введите пароль" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                        <a href="/forget" class="mt-5 text-end text-xl font-light text-gray-700">Забыли пароль?</a>
                    </div>

                    <div class="flex w-full">
                        <input type="checkbox" id="choose-me" class="peer hidden" />
                        <label for="choose-me" class="select-none cursor-pointer rounded-lg border-2 border-gray-600 bg-gray-200
                        pt-3 px-6 font-bold text-gray-600 transition-colors duration-200 text-xl ease-in-out peer-checked:bg-gray-800 peer-checked:text-white peer-checked:border-gray-200 w-full"> Запомнить аккаунт? </label>
                    </div>

                    <button
                    type="submit"
                         @click="sign_in()"
                        class="flex items-center text-xl justify-between w-full px-6 py-3  tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-500 rounded-md hover:bg-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-50">
                        Авторизоваться
                        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 rtl:-scale-x-100" viewBox="0 0 20 20" fill="currentColor">
                            <path fill-rule="evenodd"
                                d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                                clip-rule="evenodd" />
                        </svg>
                    </button>
                </form>
            </div>
        </div>
        <div class="hidden bg-cover lg:block lg:w-2/5" style="background-image: url('https://sun9-31.userapi.com/impg/HA1pFhcgIe4Hzj6PIG1cnWGcj2YwazTW8ldF8A/2ON7ThyBkto.jpg?size=900x1080&quality=95&sign=46bbc7671fb3708640b6b08a8f63d095&type=album'); background-size: 100%; background-repeat: no-repeat; background-position: center;">
        </div>
    </div>
</section>
</template>

<style scoped>
</style>