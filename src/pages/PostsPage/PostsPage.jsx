import React, { useState } from "react";
import PostsList from "../../components/UI/PostsList/PostsList";
import PostForm from "../../components/UI/PostsForm/PostsFrom";
import EditPostForm from "../../components/UI/EditPostsForm/EditPostsForm";

export default function PostsPage() {
  const [localPosts, setLocalPosts] = useState([]); // добавленные/изменённые посты
  const [editingPost, setEditingPost] = useState(null);

  function handleAdded(post) {
    // JSONPlaceholder даёт id 101 и т.д. — добавляем в локальный стейт
    setLocalPosts(prev => [post, ...prev]);
  }

  function handleEdit(post) {
    setEditingPost(post);
  }

  function handleSaved(updated) {
    setEditingPost(null);
    setLocalPosts(prev => prev.map(p => (p.id === updated.id ? updated : p)));
  }

  function handleCancel() {
    setEditingPost(null);
  }

  return (
    <div style={{ padding: 16 }}>
      <h1>Posts demo (JSONPlaceholder)</h1>
      <PostForm onAdded={handleAdded} />
      {editingPost && <EditPostForm post={editingPost} onSaved={handleSaved} onCancel={handleCancel} />}
      <PostsList onEdit={handleEdit} externalPosts={localPosts} />
    </div>
  );
}
