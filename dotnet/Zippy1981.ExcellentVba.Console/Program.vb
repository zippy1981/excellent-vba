Option Explicit

Imports System
Imports System.Runtime.InteropServices

Module Program

    Private Declare Function UUIDv1 Lib "excellent-vba" Alias "UUIDv1W" () As IntPtr
    Private Declare Function UUIDv4 Lib "excellent-vba" Alias "UUIDv4W" () As IntPtr
    Private Declare Function UUIDv6 Lib "excellent-vba" Alias "UUIDv6W" () As IntPtr
    Private Declare Function UUIDv7 Lib "excellent-vba" Alias "UUIDv7W" () As IntPtr

    Sub Main(args As String())
        Dim uuid = UUIDv7()
        Dim uuidStr = Marshal.PtrToStringBSTR(uuid)     ' This Marchall call already frees the BSTR*
        System.Console.WriteLine("v7: {0}", uuidStr)
        
    End Sub
End Module
