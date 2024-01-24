package main

import (
	"log/slog"
	"os"
)

func main() {
	// 標準出力にログを記録するハンドラーを作成
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // ログレベルを設定
	})

	// デフォルトのロガーにハンドラーをセット
	slog.SetDefault(slog.New(handler))

	// 様々なレベルのログを記録
	slog.Debug("これはデバッグメッセージです")
	slog.Info("これは情報メッセージです")
	slog.Warn("これは警告メッセージです")
	slog.Error("これはエラーメッセージです")
}

// このコードでは、slogパッケージを使って、
// 標準出力に様々なレベルのログメッセージを記録します。
// slog.NewTextHandler はテキストベースのログハンドラーを作成し、slog.SetDefault でデフォルトのロガーにセットします。
// その後、slog.Debug、slog.Info、slog.Warn、slog.Error 関数を使用して、異なるレベルのログメッセージを出力します。
