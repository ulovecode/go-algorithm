class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def preMorris(node: TreeNode, f):
    cur = node
    while cur is not None:
        mostRight: TreeNode = cur.left
        # 没有左节点
        if mostRight is not None:
            # 找到最右左节点
            while mostRight.right is not None and mostRight.right is not cur:
                mostRight = mostRight.right
            # 指回cur节点   中序
            if mostRight.right is None:
                mostRight.right = cur
                f(cur.val)
                cur = cur.left
                continue
            # 如果当前最右节点指向cur   后打印
            else:
                mostRight.right = None
        else:
            f(cur.val)
        cur = cur.right


def orderMorris(node: TreeNode, f):
    cur = node
    while cur is not None:
        mostRight: TreeNode = cur.left
        # 没有左节点
        if mostRight is not None:
            # 找到最右左节点
            while mostRight.right is not None and mostRight.right is not cur:
                mostRight = mostRight.right
            # 指回cur节点   中序
            if mostRight.right is None:
                mostRight.right = cur
                cur = cur.left
                continue
            # 如果当前最右节点指向cur   后打印
            else:
                mostRight.right = None
        else:
            pass
        f(cur.val)
        cur = cur.right


def reverseEdge(mostRight) -> TreeNode:
    pre = None
    while mostRight is not None:
        next = mostRight.right
        mostRight.right = pre
        pre = mostRight
        mostRight = next
    return pre


def printEdge(mostRight, f):
    head: TreeNode = reverseEdge(mostRight)
    while head is not None:
        f(head.val)
    reverseEdge(head)


def postmorris(node: TreeNode, f):
    cur = node
    while cur is not None:
        mostRight: TreeNode = cur.left
        # 没有左节点
        if mostRight is not None:
            # 找到最右左节点
            while mostRight.right is not None and mostRight.right is not cur:
                mostRight = mostRight.right
            # 指回cur节点   中序
            if mostRight.right is None:
                mostRight.right = cur
                cur = cur.left
                continue
            # 如果当前最右节点指向cur   后打印
            else:
                printEdge(mostRight, f)
                mostRight.right = None
        cur = cur.right
    printEdge(node, f)
