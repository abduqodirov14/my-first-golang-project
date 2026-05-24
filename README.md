# ⚡ Mening Birinchi Go (Golang) Loyiham: Telegram Web-Todo App

Salom! Bu mening Go (Golang) tilida noldan yozgan eng birinchi mustaqil loyiham. Men 16 yoshdaman va darslardan tashqari tunu kun kod yozish orqali dasturlashni o'rganyapman. Bu loyihani terminalda ishlaydigan oddiy variantdan boshlab, hozir to'liq Telegram Mini App ichida ishlaydigan darajaga olib chiqdim.

Loyiha ikkita mustaqil qismdan iborat bo'lib, o'ta tezkor ishlaydi:
1. `todo-backend` — Mening sevimli Go (Golang) tilimda yozilgan super-tezkor server qismi.
2. `todo-frontend` — HTML, CSS va toza JavaScript (Vanilla JS) yordamida kiber-qorong'u uslubda chizilgan chiroyli interfeys.

---

## 🛠️ Loyihada Men Nimalar Qildim?

* **Toza Go CRUD Logikasi:** Vazifa qo'shish, o'chirish, tahrirlash va bajarilgan deb belgilash funksiyalarini Go slice'lari orqali xotirani maksimal tejaydigan qilib yozdim.
* **XSS Xavfsizligi:** JavaScript frontend qismida `escapeHTML` funksiyasini qo'shdim. Bu foydalanuvchi inputga zararli kod yozsa ham, tizimni buzib tashlashidan (XSS hujumlaridan) asraydi.
* **CORS Global Ruxsatnomasi:** Backend serverim istalgan hostingdan keladigan so'rovlarni (fetch) xatosiz qabul qilishi uchun `*` (Allow All Origins) yulduzcha qoidasini o'rnatdim.

---


1. Terminalda backend papkasiga kiring va serverni yurgizing:
   ```bash
   cd todo-backend
   go mod init todo-backend
   go run main.go
   ```

2. Frontend qismini ishga tushirish:
   `todo-frontend/index.html` faylini shunchaki brauzerda yoki Live Server orqali oching.

---

## 🌍 Internetga Yuklash Rejam (Hosting)

* **Backend:** Go serverimni yaqin orada Render.com platformasiga tekinga yuklayman.
* **Frontend:** HTML/JS interfeysimni esa Vercel hostingiga joylashtiraman va tayyor HTTPS havolani Telegram botga ulayman!
