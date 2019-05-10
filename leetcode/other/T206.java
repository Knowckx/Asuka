package linkedlist;

public class Reverse_LinkedList_206 {

    /* 
    迭代方式，每次访问的节点，变成新链的头节点。
    */
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) { //单结点、空结点情况
            return head;
        }
        ListNode newHead = null; //newHead是新的链，最开始是空的
        ListNode p = head; //迭代的第一次是head
        
        while (p != null) {
            ListNode tmp = p.next; //mark
            p.next = newHead; //p成为新链的头
            newHead = p;  //更新newHead
            p = tmp; //下次遍历的p,前进一步
        }
        return newHead;
    }

    /*递归方式
    在反转单链表的场景下，最好的处理方式就是从最后一个节点开始搞。
    所以这里递归蛮合适的。
    */
    public ListNode reverseList2(ListNode head) {
        if (head == null || head.next == null) { //初始次和递归次都需要检查
            return head;
        }
        ListNode p = reverseList2(head.next); //直接到倒数的结点
        head.next.next = head; //4 → 5 改成了 4 ← 5 
        head.next = null; //没有这句，4 5就成了双向环
        return p;
    }

    public void test() {
        ListNode testNode = ListNode.CreatLinkedList(5);
        ListNode.PrintLinkedList(testNode);
        ListNode testResult = reverseList(testNode);
        ListNode.PrintLinkedList(testResult);

        // System.out.println(testResult);
    }

}