package title;

import java.util.Random;

public class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }

    static void PrintLinkedList(ListNode head) {
        int NodeCnt = 1;
        System.out.print("print List:\n");
        while (head != null) {
            System.out.printf("%d,%d\n", NodeCnt++, head.val);
            head = head.next;
        }
    }

    static ListNode CreatLinkedList(int length) {
        ListNode head =null;
        for (int i = 0; i < length; i++) {
            int random = (int) (Math.random()*10);
            ListNode newNode = new ListNode(random);
            newNode.next = head;
            head = newNode;
        }
        return head;
    }
}