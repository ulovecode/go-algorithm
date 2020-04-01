import math


class AvlTreeNode:
    def __init__(self, x):
        self.val = x
        self.parent: AvlTreeNode = None
        self.left: AvlTreeNode = None
        self.right: AvlTreeNode = None
        self.different = 0


def replace(head, x, y):
    if x.parent is None:
        y.parent = None
        head = y
    elif x.parent.left == x:
        set_left(x.parent, y)
    else:
        set_right(x.parent, y)
    return head


def set_left(x, l):
    x.left = l
    if l is not None:
        l.parent = x


def set_right(x, r):
    x.right = r
    if r is not None:
        r.parent = x


def set_children(x, l, r):
    set_left(x, l)
    set_right(x, r)


def left_rotate(root, x: AvlTreeNode) -> AvlTreeNode:
    y = x.right
    a = x.left
    b = y.left
    c = y.right
    root = replace(root, x, y)
    set_children(x, a, b)
    set_children(y, x, c)
    return root


def right_rotate(root, y: AvlTreeNode) -> AvlTreeNode:
    x = y.left
    a = x.left
    b = x.right
    c = y.right
    root = replace(root, y, x)
    set_children(y, b, c)
    set_children(x, a, y)
    return root


def fix_insert(head, x) -> AvlTreeNode:
    while x.parent is not None:
        p, pl, pr = x.parent, x.parent.left, x.parent.right
        d1, d2 = p.different, p.different
        if p.left == x:
            d2 = d2 - 1
        else:
            d2 = d2 + 1
        x.parent.different = d2
        if abs(d1) == 1 and abs(d2) == 0:
            return head
        elif abs(d1) is 0 and abs(d2) == 1:
            x = x.parent
        elif abs(d1) == 1 and abs(d2) == 2:
            if d2 == 2:
                if pr.different == 1:
                    p.different = 0
                    pr.different = 0
                    head = left_rotate(head, p)
                if pr.different == -1:
                    dy = pr.left.different
                    if dy == 1:
                        p.different = -1
                    else:
                        p.different = 0
                    pr.left.different = 0
                    if dy == -1:
                        pr.different = 1
                    else:
                        pr.different = 0
                    head = right_rotate(head, pr)
                    head = left_rotate(head, p)
            if d2 == -2:
                if pl.different == -1:
                    p.different = 0
                    pl.different = 0
                    head = right_rotate(head, p)
                if pl.different == 1:
                    dy = pl.right.different
                    if dy == 1:
                        pl.different = -1
                    else:
                        pl.different = 0
                    pl.right.different = 0
                    if dy == -1:
                        p.different = 1
                    else:
                        p.different = 0
                    head = left_rotate(head, pl)
                    head = right_rotate(head, p)
            break
    return head


def insert(root: AvlTreeNode, x: int) -> AvlTreeNode:
    head = root
    xNode = AvlTreeNode(x)
    p = None
    while root is not None:
        p = root
        if x < root.val:
            root = root.left
        else:
            root = root.right
    if p is None:
        return xNode
    else:
        xNode.parent = p
        if x < p.val:
            p.left = xNode
        else:
            p.right = xNode
    return fix_insert(head, xNode)


def orderPrint(root: AvlTreeNode) -> int:
    if root is None:
        return 0
    print(root.val)
    left = orderPrint(root.left)
    right = orderPrint(root.right)
    return max(left, right) + 1


if __name__ == '__main__':
    r = AvlTreeNode(0)
    for i in range(1, 1000):
        r = insert(r, i)
    print(orderPrint(r))
    print(math.log(1000,2))

