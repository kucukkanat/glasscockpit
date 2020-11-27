const webpack = require("webpack");
const path = require("path");
const CopyPlugin = require("copy-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");

const config = (env, argv) => {
  return {
    mode: env.production ? "production" : "development",
    devtool: !env.production ? "source-map" : "eval",
    entry: "./src/index.js",
    output: {
      path: path.resolve(__dirname, "dist"),
      filename: "bundle.js",
    },
    plugins: [
      new webpack.DefinePlugin({
        "process.env.NODE_ENV": JSON.stringify(process.env.NODE_ENV),
        "process.env.DEBUG": JSON.stringify(process.env.DEBUG),
      }),
      new CopyPlugin({
        patterns: [{ from: "src/assets", to: "assets" }],
      }),
      new HtmlWebpackPlugin({
        appMountId: "app",
        filename: "index.html",
        template: "src/index.html",
      }),
    ],
    module: {
      rules: [
        {
          test: /\.svg$/,
          use: "file-loader",
        },
        {
          test: /\.ts(x)?$/,
          use: [
            {
              loader: "ts-loader",
              options: {
                transpileOnly: !env.production,
              },
            },
          ],
        },
      ],
    },
    resolve: {
      extensions: [".tsx", ".ts", ".js"],
    },
    devServer: {
      contentBase: path.join(__dirname, "dist"),
      compress: true,
      port: 8080,
      host: "0.0.0.0",
    },
  };
};

module.exports = config;
