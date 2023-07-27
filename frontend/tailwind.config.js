/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'title-color':'#9DB2BF',
        'box-color': '#526D82',
        'wht-color':'#DDE6ED',
        'web-color':'#2F425B',
      },
    },
  },
  plugins: [],
}

