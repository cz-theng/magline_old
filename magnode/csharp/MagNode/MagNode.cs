using System;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
using UnityEngine;



namespace MagNode {

	internal class Node : INode{

		
		#if UNITY_STANDALONE_WIN || UNITY_EDITOR
		public const string LibName = "magnode";
		#else		
			#if UNITY_IPHONE || UNITY_XBOX360
			public const string LibName = "__Internal";
			#else
			public const string LibName = "magnode";
			#endif
		#endif

		private System.IntPtr node;

		public Node() {
			node = System.IntPtr.Zero;
			node = exp_create ();
		}

		public ErrNO Init() {
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			int rst = exp_mn_init (node);

			return (ErrNO) rst;
		}
		
		public ErrNO DeInit() {
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			int rst = exp_mn_deinit (node);
			return (ErrNO) rst;
		}
		
		public ErrNO Connect(string URL, ulong timeout) {
			Debug.Log ("Connect");
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			int rst = exp_mn_connect (node, URL, 200);
			return (ErrNO) rst;
		}
		
		public ErrNO Reconnect(ulong timeout) {
			Debug.Log ("Reconnect");
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			return ErrNO.Success;
		}
		
		public ErrNO Send(byte[] data,uint length,ulong timeout) {
			Debug.Log ("Send message" +data);
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}

			int rst = exp_mn_send (node, data, (int)length,(int) timeout);

			return (ErrNO) rst;
		}
		
		public ErrNO Recv(ref byte[] buf,ref uint length,ulong timeout) {
			Debug.Log ("Recv");
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			int len =(int) length;
			int rst = exp_mn_recv (node, buf,ref len, (int)timeout);

			Debug.Log ("exp_mn_recv rst is "+rst);
			if (rst == 0) {
				length = (uint)len;
			}
			return ErrNO.Success;
		}
		
		public ErrNO Close() {
			if (System.IntPtr.Zero == node) {
				return ErrNO.Nil;
			}
			int rst = exp_mn_close (node);
			return ErrNO.Success;
		}

		#region DllImport 
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern System.IntPtr exp_create();
		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_init(System.IntPtr node);

		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_deinit(System.IntPtr node);

		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_connect(System.IntPtr node, [MarshalAs(UnmanagedType.LPArray)] string key, int timeout);

		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_send(System.IntPtr node, [MarshalAs(UnmanagedType.LPArray)] byte[] buf, int lenght, int timeout);

		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_recv(System.IntPtr node, [MarshalAs(UnmanagedType.LPArray)]byte[] buf, ref int size, int timeout);

		
		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_close(System.IntPtr node);

		[DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
		private static extern int exp_mn_destory(System.IntPtr node);
		#endregion

	}
}
