<template>
<div class="main relative h-screen">
    <div class="max-w-7xl absolute top-20 left-1/4 w-1/2  mx-auto">
        <section class=" p-6  bg-indigo-600 rounded-md shadow-md dark:bg-gray-800 ">
            <h1 class="text-2xl font-bold text-white capitalize dark:text-white">Редактирование профиля</h1>
            <form>
                <div class="grid grid-cols-1 gap-6 mt-4 sm:grid-cols-2">
                    <div class="row-span-2">
                        <label class="block text-xl font-medium text-white">
                        Аватарка
                    </label>
                    <div class="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
                        <div class="space-y-1 text-center">
                        <svg class="mx-auto h-12 w-12 text-white" stroke="currentColor" fill="none" viewBox="0 0 48 48" aria-hidden="true">
                            <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                        </svg>
                            <div class="flex text-sm text-gray-600">
                            <label for="file-upload" class="relative cursor-pointer bg-white rounded-md font-medium text-indigo-600 hover:text-indigo-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-indigo-500">
                                <span class="p-2">Upload a file</span>
                                <div class="form-group">
                                    <div class="col-sm-8">
                                        <input type="file" class="sr-only">
                                    </div>
                                </div>
                            </label>
                            <p class="pl-1 text-white">or drag and drop</p>
                            </div>
                            <p class="text-xs text-white">
                            PNG, JPG, GIF up to 10MB
                            </p>
                        </div>

                    </div>
                    </div>

                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="username">Логин</label>
                        <input type="text" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                    </div>

                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="emailAddress">Почта</label>
                        <input type="email" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                            <!-- <div>Currently waiting confirmation for: <%= resource.unconfirmed_email %></div> -->
                    </div>

                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="password">Пароль</label>
                        <input type="password" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                        <em class="text-gray-200">6 characters minimum</em>
                    </div>

                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="passwordConfirmation">Подтверждение пароля</label>
                        <input type="password" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                    </div>
                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="passwordConfirmation">Текущий пароль</label>
                        <input type="password" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                    </div>
                    <div>
                        <label class="text-white text-xl dark:text-gray-200" for="passwordConfirmation">Город</label>
                        <input type="select" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring">
                    </div>
                    <div class="col-span-2">
                        <label class="text-white text-xl dark:text-gray-200" for="passwordConfirmation">О себе</label>
                        <textarea maxlength="250"  style="resize:none" class=" block  w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-300 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring"/>
                    </div>

                </div>

                <div class="flex justify-end mt-6">
                    <button type="submit" class="px-6  py-2 leading-5 text-white transition-colors duration-200 transform bg-pink-500 rounded-md hover:bg-pink-700 focus:outline-none focus:bg-gray-600">Сохранить</button>
                </div>
            </form>
        </section>
    </div>
</div>

</template>
<style scoped>
    .main {
        background-image: url('https://sun9-54.userapi.com/impg/VByb1K2lGE3rvesgj9tL6NgKM4FS1LDvtyZWbg/S3mtN3Udnic.jpg?size=2048x1463&quality=95&sign=99a3a94f6c12c70fe2ec894e32d17928&type=album'); 
        background-size: 100% 100%; 
        background-repeat: no-repeat; 
        background-position: center;
    }
</style>