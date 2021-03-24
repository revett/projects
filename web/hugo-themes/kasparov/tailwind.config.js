module.exports = {
  corePlugins: {
    preflight: false,
  },
  darkMode: false,
  plugins: [],
  purge: ["../../**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        serif: ["Source Sans Pro", "sans-serif"],
      },
    },
  },
  variants: {
    extend: {},
  },
};
