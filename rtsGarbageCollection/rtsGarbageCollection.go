package rtsGarbageCollection

import (
    "runtime"
    "runtime/debug"
)

// 強制ガベージコレクタの実行
func GC() {
    runtime.GC()
}

// golangのガベージコレクタを動かさない
func GcOff() {
    GC()
    debug.SetGCPercent(-1)
}

// golangのガベージコレクタを動かす
func GcOn() {
    GC()
    debug.SetGCPercent(100)
}

// 削除されたデータが500MBならガベージコレクタを
func GcTimer(){
    var mem runtime.MemStats

    runtime.ReadMemStats(&mem)

    if mem.HeapAlloc > 500 << 20 { // 500MB
        GC()
    }
}