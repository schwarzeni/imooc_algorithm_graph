## 忘记的概念

### 图的基本表示

- 联通分量（2-2 8:13）
- 树是一种无环图（2-2 10:26）
- 联通的无环图是树（2-2 11:27）
- 联通图的生成树包含所有顶点（2-2 12:38）
- 通过删边可以使原本的联通图不在联通（2-2 14:40）
- 只有联通图才有生成树（2-2 15:20）
- 邻接表空间复杂度表示 O(V+E)，不能用 O(E)，当 E=0 的时候就不对了（2-7 05:11）
- 邻接表查询两点是否相邻优化：使用 HashSet（哈希表） 或者 TreeSet（红黑树）（2-7 13:35）
- 由于 Golang 中没有内置的 TreeSet 红黑树，这里就使用 map （哈希表）了 :/
- 可以使用深度优先遍历
    - 验证图是否联通、是否有环
    - 二分图检测
    - 寻找图中的桥
    - 寻找图中的割点
    - 哈密尔顿路径
    - 拓扑排序

---

### 深度优先遍历的应用

1. 求联通分量的个数 (4-1)

```text
visited[0 .. V] = false;

for (int v = 0; v < V; v ++) {
    if (!visited[v]) {
        dfs(v);
        // connected component
        cccount ++
    }
}
```

2. 求联通分量中包含的顶点 (4-2)

`visited` 数组记录每一个联通分量的 ID ，属于同一个联通分量的 vertex 有相同的 ID

```text
visited[0 .. V] = -1;

for (int v = 0; v < V; v ++) {
    if (!visited[v]) { // set ID
        dfs(v, ccid);
        ccid ++
    }
}
```

3. 求两个点是否在同一个联通分量重 (4-3)
4. 求两点之间的路径 (4-4, 单源路径问题)
5. 二分图检测
    - 二分图：定点 V 可以分成不想交的两个部分，所有边的两个端点隶属于不同部分

```text
0 - 1
0 - 2
3 - 1
3 - 2
4 - 1
6 - 2
6 - 5
```

---

### 关于广度优先遍历

使用广度优先遍历求得的路径即为最短路径，但是图必须为**无权图**

BFS 和 DSF 拥有一套相似的逻辑框架

```text
visited[0 ... V-1] = false

for (int v = 0; v < V; v ++) {
    if (!visited[v]) {
        search(v)
    }
}

func search(int s) {
    x.add(s)
    visited[s] = true
    while(!x.isEmpty()) {
        int v = x.remove();
        
        for (int w : G.adj(v)) {
            if (!visited[w]) {
                x.add(w)
                visited[w] = true
            }
        }

    }
}
```

---

## 第六章 floodfill

- [695](leetcode/l-695)
- [200](leetcode/l-200)
- [1020](leetcode/l-1020) 逆向思维，排除法
- [130](leetcode/l-130) 逆向思维，排除法
- [733](leetcode/l-733)   图像渲染 
- [1034](leetcode/l-1034) 边框着色
- [529](leetcode/l-529) 扫雷
- [827](leetcode/l-827) 最大人工岛 两步法解题

---

## 第七章 BFS

抽象状态和状态转移

- [1091](leetcode/l-1091) 二进制矩阵中的最短路径
- [752](leetcode/l-752)  打开转盘锁，使用图论建模
- [7-4一道智力题](traverse_graph/7-4-IQ-test) 倒水问题
- [7-4二道智力题](traverse_graph/7-4-IQ-test) 农夫运狼羊菜
