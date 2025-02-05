import { request } from '@/utils'

export default {
  // 创建文件/文件夹
  create: (path: string, dir: boolean): any => request.post('/file/create', { path, dir }),
  // 获取文件内容
  content: (path: string): any => request.get('/file/content', { params: { path } }),
  // 保存文件
  save: (path: string, content: string): any => request.post('/file/save', { path, content }),
  // 删除文件
  delete: (path: string): any => request.post('/file/delete', { path }),
  // 上传文件
  upload: (path: string, formData: FormData, onProgress: any): any => {
    formData.append('path', path)
    return request.post('/file/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (progressEvent: any) => {
        onProgress({ percent: Math.ceil((progressEvent.loaded / progressEvent.total) * 100) })
      }
    })
  },
  // 检查文件是否存在
  exist: (paths: string[]): any => request.post('/file/exist', paths),
  // 移动文件
  move: (paths: any[]): any => request.post('/file/move', paths),
  // 复制文件
  copy: (paths: any[]): any => request.post('/file/copy', paths),
  // 远程下载
  remoteDownload: (path: string, url: string): any =>
    request.post('/file/remoteDownload', { path, url }),
  // 获取文件信息
  info: (path: string): any => request.get('/file/info', { params: { path } }),
  // 修改文件权限
  permission: (path: string, mode: string, owner: string, group: string): any =>
    request.post('/file/permission', { path, mode, owner, group }),
  // 压缩文件
  compress: (dir: string, paths: string[], file: string): any =>
    request.post('/file/compress', { dir, paths, file }),
  // 解压文件
  unCompress: (file: string, path: string): any => request.post('/file/unCompress', { file, path }),
  // 搜索文件
  search: (path: string, keyword: string, sub: boolean, page: number, limit: number): any =>
    request.get('/file/search', { params: { path, keyword, sub, page, limit } }),
  // 获取文件列表
  list: (path: string, page: number, limit: number, sort: string): any =>
    request.get('/file/list', { params: { path, page, limit, sort } })
}
