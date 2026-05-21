<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const roomId = route.params.id

const canvasRef = ref(null)
const ctx = ref(null)
const ws = ref(null)
const color = ref('#000000')

let isDrawing = false
let lastX = 0
let lastY = 0

// إعداد الروابط ديناميكياً لتشير إلى السيرفر الخاص بك عند الرفع
// استبدل 'your-backend-domain.com' بنطاق السيرفر الفعلي الخاص بك لاحقاً
const IS_PROD = window.location.hostname !== 'localhost'
const BACKEND_URL = IS_PROD ? 'your-backend-domain.com' : 'localhost:8080'

const HTTP_PROTOCOL = window.location.protocol === 'https:' ? 'https://' : 'http://'
const WS_PROTOCOL = window.location.protocol === 'https:' ? 'wss://' : 'ws://'

onMounted(async () => {
  console.log('Board mounted, room:', roomId)
  
  const canvas = canvasRef.value
  if (!canvas) {
    console.error('Canvas not found!')
    return
  }
  
  // تعيين الحجم الفعلي للـ Canvas ليطابق مساحة العرض
  canvas.width = canvas.parentElement.clientWidth
  canvas.height = window.innerHeight - 60
  
  ctx.value = canvas.getContext('2d')
  ctx.value.lineWidth = 3
  ctx.value.lineCap = 'round'

  // جلب التاريخ برابط ديناميكي
  try {
    const res = await fetch(`${HTTP_PROTOCOL}${BACKEND_URL}/api/rooms/${roomId}/history`)
    if (res.ok) {
      const history = await res.json()
      history.forEach(stroke => {
        // حماية الكود: التأكد إذا كان العنصر نصاً يحتاج لتحليل أم كائناً جاهزاً
        const data = typeof stroke === 'string' ? JSON.parse(stroke) : stroke
        drawLine(data.x0, data.y0, data.x1, data.y1, data.color, false)
      })
    }
  } catch (e) {
    console.log('No history or error:', e)
  }

  // WebSocket برابط ديناميكي
  ws.value = new WebSocket(`${WS_PROTOCOL}${BACKEND_URL}/ws?room=${roomId}`)
  
  ws.value.onopen = () => console.log('WebSocket connected!')
  ws.value.onerror = (e) => console.error('WebSocket error:', e)
  
  ws.value.onmessage = (event) => {
    const data = JSON.parse(event.data)
    drawLine(data.x0, data.y0, data.x1, data.y1, data.color, false)
  }
})

onUnmounted(() => {
  if (ws.value) ws.value.close()
})

function startDrawing(e) {
  isDrawing = true
  lastX = e.offsetX
  lastY = e.offsetY
}

function draw(e) {
  if (!isDrawing) return
  
  const currentX = e.offsetX
  const currentY = e.offsetY

  drawLine(lastX, lastY, currentX, currentY, color.value, true)

  lastX = currentX
  lastY = currentY
}

function stopDrawing() {
  isDrawing = false
}

function drawLine(x0, y0, x1, y1, strokeColor, sendToWS) {
  const c = ctx.value
  if (!c) return // حماية إضافية في حال لم يتم تحميل السياق بعد
  
  c.beginPath()
  c.moveTo(x0, y0)
  c.lineTo(x1, y1)
  c.strokeStyle = strokeColor
  c.stroke()

  if (sendToWS && ws.value?.readyState === WebSocket.OPEN) {
    const payload = JSON.stringify({ x0, y0, x1, y1, color: strokeColor })
    ws.value.send(payload)
  }
}
</script>

<template>
  <div class="h-screen flex flex-col overflow-hidden">
    <!-- Toolbar -->
    <div class="h-[60px] bg-white border-b flex items-center px-4 gap-4 z-10">
      <button @click="router.push('/')" class="px-3 py-1 bg-gray-200 rounded hover:bg-gray-300">
        العودة
      </button>
      <span class="font-bold border-l pl-4 border-gray-300">Room: {{ roomId }}</span>
      
      <span>اختر اللون:</span>
      <input type="color" v-model="color" class="w-8 h-8 rounded cursor-pointer border-0" />
    </div>

    <!-- Canvas مع إضافة flex-1 و w-full لحل مشكلة الاختفاء البصري -->
    <canvas 
      ref="canvasRef"
      @mousedown="startDrawing"
      @mousemove="draw"
      @mouseup="stopDrawing"
      @mouseleave="stopDrawing"
      class="bg-gray-50 cursor-crosshair flex-1 w-full"
    ></canvas>
  </div>
</template>