<?php
class Node {
    public $key;
    public $value;
    public $pre;
    public $next;

    public function __construct($key, $value) {
        $this->key = $key;
        $this->value = $value;
    }

}

class LRUcache {
    public  $map = array();
    protected $head;
    protected $end;
    protected $capacity;

    function __construct($capacity) {
        $this->capacity = $capacity;
    }

    //get() 当数据项被查到时，将此数据项移动到链表头部，时间复杂度O(1)
    function get($key) {
        if (isset($this->map[$key])) {
            $node = $this->map[$key];
            //移除节点
            $this->remove($node);
            //插入链表头部
            $this->setHeader($node);
            return $node->value;
        }
        return -1;
    }

    /**
     * 移除节点
     * @param Node $node
     */
    function remove(Node $node) {
        //存在前驱，将前驱节点的next指针指向后一节点
        //不存在，则将后继节点赋给头指针
        if ($node->pre != null) {
            $node->pre->next = $node->next;
        } else {
            $this->head = $node->next;
        }
        //存在后继节点，则将后继节点的pre指针指向前一节点
        //不存在，则将后继节点赋给尾指针
        if ($node->next != null) {
            $node->next->pre = $node->pre;
        } else {
            $this->end = $node->pre;
        }
    }

    /**
     * 设置头节点
     * (1) 取出头节点，将头节点的 pre指针，指向新节点
     * (2) 将head指针，指向新节点
     * (3) 如果达到哈希表容量，移除最后一个节点
     * (4) 
     */
    protected function setHeader(Node $node)
    {
        $node->next = $this->head;
        $node->pre = null;

        if ($this->head != null) {
            $this->head->pre = $node;       
        }
        $this->head = $node;
        if ($this->end == null) {
            $this->end = $node;
        }
    }

    /**
     * 添加节点，如果key存在，获取旧节点，更新值，移除旧节点，并设置到头节点
     * key不存在，判断哈希表是否达到容量
     * 达到容量，销毁哈希表尾节点对应元素，移除尾节点，设置头节点为新节点
     * 不达到容量,设置头节点为新节点
     */
    public function set($key, $value) {
        if (isset($this->map[$key])) {
            $old = $this->map[$key];
            $old->value = $value;
            $this->remove($old);
            $this->setHeader($old);
        } else {
            $node = new Node($key, $value);
            if (count($this->map) >= $this->capacity) {
                unset($this->map[$this->end->key]);
                $this->remove($this->end);
            }
            $this->setHeader($node);
            $this->map[$key] = $node;
        }
    }
}

$cache = new LRUCache(2);
$cache->set(1, 1);
$cache->set(2, 2);
$cache->get(1);
$cache->set(3, 3);
$res = $cache->get(2);
var_dump($res);
var_dump($cache->map);