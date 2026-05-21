<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './store/store'

const username = ref('')
const password = ref('')
const auth = useAuthStore()
const router = useRouter()

async function handleAuth(type) {
  try {
    const res = await fetch(`http://localhost:8080/api/${type}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'  // ← هذا السطر هو الحل!
      },
      body: JSON.stringify({ 
        username: username.value, 
        password: password.value 
      })
    })
    
    if (res.ok && type === 'login') {
      const data = await res.json()
      auth.setToken(data.token)
      router.push('/board/room1')  // تحويل مباشر للغرفة
    } else if (res.ok && type === 'register') {
      alert("تم التسجيل! اضغط Login الآن")
    } else {
      const errorText = await res.text()
      alert("خطأ: " + errorText)
    }
  } catch (error) {
    alert("لا يمكن الاتصال بالسيرفر. تأكد أنه يعمل على المنفذ 8080")
  }
}
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-96">
      
      <div v-if="!auth.token" class="flex flex-col gap-4">
        <h2 class="text-xl font-bold text-center">أهلاً بك في السبورة</h2>
        <input v-model="username" placeholder="اسم المستخدم" class="border p-2 rounded" />
        <input v-model="password" type="password" placeholder="كلمة المرور" class="border p-2 rounded" />
        <div class="flex gap-2">
          <button @click="handleAuth('login')" class="bg-blue-500 text-white w-full py-2 rounded">دخول</button>
          <button @click="handleAuth('register')" class="bg-green-500 text-white w-full py-2 rounded">تسجيل</button>
        </div>
      </div>

      <div v-else class="text-center flex flex-col gap-4">
        <h2 class="text-xl font-bold">مرحباً بك!</h2>
        <button @click="router.push('/board/room1')" class="bg-blue-600 text-white py-3 rounded text-lg font-bold">
          ➕ إنشاء غرفة جديدة
        </button>
        <button @click="auth.logout()" class="text-red-500 underline mt-4">تسجيل الخروج</button>
      </div>

    </div>
  </div>
</template>