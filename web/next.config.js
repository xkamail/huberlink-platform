const withPWA = require('next-pwa')({
    dest: "public",
    register: true,
    skipWaiting: true,
})

/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        appDir: true,
        // runtime: "experimental-edge"
    },
    reactStrictMode: true,
    swcMinify: true,
};

module.exports = withPWA(nextConfig);
