import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  build: {
    lib: {
      entry: "src/index.ts",
      name: "components",
      formats: ["es", "cjs"]
    },
    rollupOptions: {
      external: ["react", "react-dom"]
    }
  }
});
