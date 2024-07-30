<script setup>
</script>


<template>
   <section class="bg-white w-full dark:bg-white height-[100vh]" style="height:100vh">
    <div class="flex h-full bg-white">
        <div class="main hidden bg-fill lg:block lg:w-2/5"></div>
        <div class="flex items-center w-full md:w-1/2 p-8 bg-white mx-auto lg:px-12 ">
            <div class="w-full mt-24">
                <h1 class="text-4xl mt-24 md:mt-14 font-semibold tracking-wider text-gray-800 capitalize ">
                    Создай аккаунт прямо сейчас.
                </h1>

                <p class="mt-4 text-2xl text-gray-500 dark:text-gray-400">
                    Благодаря RedNews вы сможете получать самые актуальные новости на данный момент. Если же у вас уже имеется аккаунт - <a href="/sign_in" class="text-blue-500">Авторизация</a> 
                </p>


                <div class="grid w-full  sn:grid-cols-1 gap-6 mt-8 md:grid-cols-2 lg:grid-cols-2">
                    <div class="mt-2">
                        <label class="block mb-2 text-xl text-gray-600 ">Логин</label>
                        <input type="text" placeholder="Введите логин" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                    </div>

                    <div class="mt-2">
                        <label class="block mb-2 text-xl text-gray-600 ">О себе</label>
                        <input type="text" placeholder="Введите о себе" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                    </div>

                    <div class="mt-2">
                        <label class="block mb-2 text-xl text-gray-600 ">Почта</label>
                        <input type="email" placeholder="Введите почту" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                    </div>

                    <div class="mt-2">
                        <label class="block mb-2 text-xl text-gray-600">Пароль</label>
                        <input type="password"  placeholder="Введите пароль" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                        <em class="text-sm text-gray-500">(6 знаков минимум)</em>
                    </div>

                    <div class="mt-2">
                        <label class="block mb-2 text-xl text-gray-600 ">Подтверждение пароля</label>
                        <input type="password" placeholder="Введите снова пароль" class="block w-full px-3 py-3 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40">
                    </div>

                    <div class="md:pt-9 md:mb-10 mb-0 lg:pt-11">
                      <button
                        style="height:50px"
                        class="flex text-xl items-center justify-between w-full px-6 py-3 tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-500 rounded-md hover:bg-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-50">
                        Войти
                        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 rtl:-scale-x-100" viewBox="0 0 20 20" fill="currentColor">
                            <path fill-rule="evenodd"
                                d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                                clip-rule="evenodd" />
                        </svg>
                      </button>
                    </div>

                </div>
            </div>
        </div>
    </div>
</section>
</template>

<style scoped>
.main {
    background-image: url('../assets/SignUp.jpg'); background-size: 100%; background-repeat: no-repeat; background-position: bottom;
}
</style>