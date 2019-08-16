<?php

// LeetCode 146 题
Class LRUCache 
{

  // 缓存数组 左:新 -> 右:旧
  public $cacheList = [];

  // 缓存数组长度
  private $cacheLength; 

  public function __construct($lenth) {
      $this->cacheLength = $lenth;
  }

  // 新增
  public function put($key, $value) {

    unset($this->cacheList[$key]);

    // 缓存池已满 -> 移除最右
    if ($this->checkLength()) {
        array_pop($this->cacheList);
    }

    // 写入最左
    $this->cacheList = [$key => $value] + $this->cacheList;
  }

  // 获取
  public function get($key) {

    if (!isset($this->cacheList[$key])) {
      return -1;
    }

    // 更新位置到最右
    $value = $this->cacheList[$key];
    unset($this->cacheList[$key]);
    $this->cacheList = [$key => $value] + $this->cacheList;

    return $value;

  }

  // 缓存池是否已满
  private function checkLength() {
    return count($this->cacheList) >= $this->cacheLength;
  }
}



/*

LRUCache cache = new LRUCache( 2  );  // 缓存容量

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4

*/
