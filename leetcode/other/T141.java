package linkedlist;


public class Cycle_141 {

    /* 
    两个指针，步长分别为一步和两步。是否会相遇。
    */
    public boolean hasCycle(ListNode head) {
        if (head == null || head.next == null) { //单结点、空结点情况
            return false;
        }

        ListNode p1 = head;
        ListNode p2 = head;

        //p2结点有可能已是最后一个。这样调p2.next.next就是空指针了。
        while (p2.next != null && p2.next.next != null) {
            p1 = p1.next;
            p2 = p2.next.next;
            if (p1 == p2) {
                return true;
            }
        }
        return false;
    }

    /* 每遍历一个，就把该节点指向自己。 
    假如在后面的遍历中发现有指向自己的。那就是被访问过的，就是环。
    需要额外的1空间。
    */
    public boolean hasCycle2(ListNode head) {
        if (head == null || head.next == null) { //单结点、空结点情况
            return false;
        }
        if (head.next == head) {
            return true;
        }
        ListNode p1 = head.next;
        head.next = head;     //mark已访问
        boolean isCycle = hasCycle2(p1);
        return isCycle;
    }

}