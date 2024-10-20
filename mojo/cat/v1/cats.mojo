@entity("Cats")
interface Cats {
    /// 创建cat信息
    @http.post("/tnr/v1/cats")
    create_cat(cat: Cat @1)-> Cat

    /// 更新cat信息
    @http.put("/tnr/v1/cats/{cat.id}")
    update_cat(cat: Cat @1)-> Cat

    /// 查询cat信息
    @http.get("/tnr/v1/cats/{id}")
    get_cat(id: String @1)-> Cat

    /// 查询cat列表
    @http.get("/tnr/v1/cats")
    list_cats()-> [Cat]

    /// 批量获取cat信息
    @http.get("/tnr/v1/cats:batch-get")
    batch_get_cat(ids: [String] @1)-> [Cat]

    /// 删除cat信息
    @http.delete("/tnr/v1/cats/{id}")
    delete_cat(id: String @1)
}
